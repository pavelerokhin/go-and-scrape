import spacy
import sqliteconnector as s


nlp = spacy.load('ru_core_news_sm')
for text in s.get_texts("mediums"):
    doc = nlp(text)
    for ent in doc.ents:
        print(ent)