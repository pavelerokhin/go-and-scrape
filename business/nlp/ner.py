import spacy

import cleaning as c
import sqliteconnector as s


nlp = spacy.load('ru_core_news_sm')
for text in s.get_texts("mediums"):
    t = c.spoil_ebala(text)
    print(t[:100], "...")
    doc = nlp(t)
    for ent in doc.ents:
        print(ent)