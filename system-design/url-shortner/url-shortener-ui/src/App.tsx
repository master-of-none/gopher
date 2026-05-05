import React, { useState } from "react";
import { shortenURl } from "./api";

function App() {
  const [url, setUrl] = useState("");
  const [expiry, setExpiry] = useState("");
  const [shortUrl, setShortUrl] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const [history, setHistory] = useState<string[]>(() => {
    return JSON.parse(localStorage.getItem("history") || "[]");
  });

  const saveUrl = (url: string) => {
    const existing = JSON.parse(localStorage.getItem("history") || "[]");
    const updated = [url, ...existing.filter((u: string) => u != url)].slice(
      0,
      5,
    );
    localStorage.setItem("history", JSON.stringify(updated));
    setHistory(updated);
  };
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setShortUrl("");
    setLoading(true);

    try {
      const data = await shortenURl(
        { url },
        expiry ? Number(expiry) : undefined,
      );
      setShortUrl(data.short_url);
      saveUrl(url);
    } catch (e: unknown) {
      setError(e instanceof Error ? e.message : "Unknown Error");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center">
      <div className="w-fit max-w-md bg-white rounded-2xl shadow-lg p-8 space-y-6">
        <h1 className="text-2xl font-semibold text-center">URL Shortener</h1>
        <form onSubmit={handleSubmit} className="space-y-4">
          <input
            type="text"
            placeholder="Enter URL"
            value={url}
            onChange={(e) => setUrl(e.target.value)}
            className="w-full border rounded-lg p-3 focus:outline-none focus:ring-2 focus:ring-blue-200"
          />
          <input
            type="number"
            placeholder="Expiry (minutes)"
            value={expiry}
            onChange={(e) => setExpiry(e.target.value)}
            className="w-full border rounded-lg p-3 focus:outline-none focus:ring-2 focus:ring-blue-300"
          />
          {url && history.length > 0 && (
            <div className="bg-white border rounded mt-2 shadow max-h-40 overflow-y-auto">
              {history.map((item, i) => (
                <div
                  key={i}
                  onClick={() => setUrl(item)}
                  className="p-2 hover:bg-gray-100 cursor-pointer text-sm"
                >
                  {item}
                </div>
              ))}
            </div>
          )}
          <button
            type="submit"
            disabled={loading}
            className="w-full bg-blue-700 text-white p-3 rounded-lg hover:bg-blue-800 transition disabled:opacity-40"
          >
            {loading ? "Shortening" : "Shorten"}
          </button>
        </form>
        {shortUrl && (
          <div className="bg-green-200 p-4 rounded-lg text-center space-y-2">
            <p className="text-lg text-gray-600"> Short URL </p>
            <a
              href={shortUrl}
              target="_blank"
              rel="noreferrer"
              className="text-blue-800 font-medium break-all"
            >
              {shortUrl}
            </a>
          </div>
        )}
        {error && (
          <div className="bg-red-500 text-red-200 p-4 rounded-lg text-center">
            {error}
          </div>
        )}
      </div>
    </div>
  );
}

export default App;
