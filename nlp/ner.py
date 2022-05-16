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


def levenshtein_distance(w1, w2):
    # for all i and j, d[i, j] will hold the Levenshtein distance between
    # the first i characters of s and the first j characters of t
    n = len(w1)
    m = len(w2)
    d = [[0]*n]*m

    # source prefixes can be transformed into empty string by dropping all characters
    for i in range(n):
        d[i][0] = i

    # target prefixes can be reached from empty source prefix by inserting every character
    for j in range(n):
        d[0][j] = j

    for j in range(n):
        for i in range(m):
            if w1[i] == w2[j]:
                substitutionCost = 0
            else:
                substitutionCost = 1

            d[i][j] = min(d[i - 1][j] + 1, # deletion
            d[i][j - 1] + 1, # insertion
            d[i - 1][j - 1] + substitutionCost) # substitution

    return d[m][n]


def levenshtein_distance_filter(w1, w2, limit=2):
    # returns true if distance between w1 and w2 is bigger then limit
    return levenshtein_distance(w1, w2) > limit
