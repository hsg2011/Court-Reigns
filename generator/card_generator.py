
import os
import json
import re # for cleaning up the JSON output
from openai import OpenAI   # the new unified client for OpenAI + Gemini

def clean_json(s: str) -> str:
    return re.sub(r"^```(?:\w+)?\s*|\s*```$", "", s, flags=re.MULTILINE)

def load_api_key(path="gemini-key.txt"):
    with open(path, "r") as f:
        return f.read().strip()

def generate_cards(client, n=20) -> str:
    system_prompt = (
        "You are a generator for a basketball-team management game. "
        "Produce exactly {n} decision cards in a JSON array, where each card is an object with:\n"
        "  - \"text\": a one-sentence scenario string,\n"
        "  - \"left\": {{\"finances\": int, \"morale\": int, \"fitness\": int, \"fans\": int}},\n"
        "  - \"right\": same shape.\n"
        "All deltas must be between -50 and +50. Return ONLY the JSON array."
    ).format(n=n)

    response = client.chat.completions.create(
        model="gemini-2.0-flash",           # or whichever AI model you have access to
        messages=[
            {"role":"system", "content": system_prompt},
            {"role":"user",   "content": "Please generate the cards now."},
        ],
        temperature=0.7,
        n=1,
    )
    # The content field holds the JSON text
    return response.choices[0].message.content

def main():
    # 1) Read your Gemini key
    api_key = load_api_key("gemini-key.txt")

    # 2) Instantiate the client, pointing at the Generative Language API
    client = OpenAI(
        api_key=api_key,
        base_url="https://generativelanguage.googleapis.com/v1beta/openai/"
    )

    # 3) Generate cards (change the count if you like)
    print("Generating cards via Gemini…")
    cards_json = generate_cards(client, n=5)
    cleaned = clean_json(cards_json)

    # 4) Validate or pretty-print :
    try:
        cards = json.loads(cleaned)
    except json.JSONDecodeError as e:
        print("Failed to parse JSON from Gemini:", e)
        print("---- raw response ----")
        print(cleaned)
        return

    # 5) Write back into cards.json
    with open("cards.json", "w") as f:
        json.dump(cards, f, indent=2)
    print(f"Wrote {len(cards)} cards into cards.json")

if __name__ == "__main__":
    main()
