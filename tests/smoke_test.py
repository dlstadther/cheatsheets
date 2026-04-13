BASE_URL = "http://localhost:8000"


def test_root_returns_200(page):
    response = page.goto(f"{BASE_URL}/")
    assert response.status == 200


def test_nav_links_load_without_error(page):
    page.goto(f"{BASE_URL}/")

    nav_links = page.locator(".md-nav a[href]").all()
    hrefs = set()
    for link in nav_links:
        href = link.get_attribute("href")
        if href and not href.startswith("http") and not href.startswith("#"):
            hrefs.add(href.split("#")[0])

    assert hrefs, "No nav links found — check .md-nav selector against the built site"

    for href in hrefs:
        url = f"{BASE_URL}{href}" if href.startswith("/") else f"{BASE_URL}/{href}"
        response = page.goto(url)
        assert response.status == 200, f"Nav link {href} returned HTTP {response.status}"
