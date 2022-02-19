import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

let timelineData = [
  {
    text: 'Wrote my first blog post ever on Medium',
    date: 'March 03 2017',
    category: {
      tag: 'medium',
      color: '#018f69'
    },
    link: {
      url:
        'https://medium.com/@popflorin1705/javascript-coding-challenge-1-6d9c712963d2',
      text: 'Read more'
    }
  },
  {
    text: 'Wrote my first blog post dsfsdfdf on Medium',
    date: 'March 03 2018',
    category: {
      tag: 'medium',
      color: '#018f69'
    },
    link: {
      url:
        'https://medium.com/@popflorin1705/javascript-coding-challenge-1-6d9c712963d2',
      text: 'Read more'
    }
  }
];

const Timeline = () =>
  timelineData.length > 0 && (
    <div className="timeline-container">
      {timelineData.map((data, idx) => (
        <TimelineItem data={data} key={idx} />
      ))}
    </div>
  );

const TimelineItem = ({ data }) => (
  <div className="timeline-item">
    <div className="timeline-item-content">
      <span className="tag" style={{ background: data.category.color }}>
        {data.category.tag}
      </span>
      <time>{data.date}</time>
      <p>{data.text}</p>
      {data.link && (
        <a
          href={data.link.url}
          target="_blank"
          rel="noopener noreferrer"
        >
          {data.link.text}
        </a>
      )}
      <span className="circle" />
    </div>
  </div>
);

export default Timeline;
