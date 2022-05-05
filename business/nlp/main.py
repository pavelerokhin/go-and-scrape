import ner as n
import sqliteconnector as s


def ner_worker(db_name, ids):
    articles = s.get_texts(db_name, ids)
    articles_lemmas = n.articles_ner_and_count(articles)
    if len(articles_lemmas):
        s.persist_articles_lemmas(articles_lemmas)
