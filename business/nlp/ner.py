import spacy

import cleaning as c

nlp = spacy.load('ru_core_news_sm')


def articles_ner_and_count(articles):
    articles_lemmas = {}
    for (id, text) in articles:
        t = c.spoil_ebala(text)
        doc = nlp(t)
        lemmas = {}
        for ent in doc.ents:
            key = f"{ent.lemma_}_{ent.label_}"
            if lemmas.get(key):
                lemmas[key] += 1
            else:
                lemmas[key] = 1
        articles_lemmas[id] = lemmas

    return articles_lemmas
