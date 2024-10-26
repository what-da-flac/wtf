export function                      HomeSection() {
  return (
    <div className="publux-container">
      <div className="publux-heading">
        <h1>Cinito</h1>
        <div className="description">
          The stupid web application that allows users to share movie reviews.
        </div>
      </div>
      <div className="flex-lists">
        <div className="form-wrapper">
          <h2>How does it work?</h2>
          <ol>
            <li>Login with your Google account.</li>
            <li>
              Administrator will review your information and send you an
              invitation to join the community, otherwise an apology if your
              request was rejected.
            </li>
            <li>Browse movie list and start reviewing your favorites.</li>
          </ol>
        </div>
        <div className="form-wrapper">
          <h2>Motivation</h2>
          <ul>
            <li>It is difficult to get movie reviews from indie movies.</li>
            <li>
              Avoid watching doggy-shit movies, like those on streaming
              services.
            </li>
            <li>Meet new fellows with similar movie taste.</li>
          </ul>
        </div>
        <div className="form-wrapper">
          <h2>Notes</h2>
          <ul>
            <li>Cinito is not a paid application.</li>
            <li>
              Founder covers the costs and development because it is a hobby, so
              don't expect paid support :).
            </li>
          </ul>
        </div>
      </div>
    </div>
  );
}
