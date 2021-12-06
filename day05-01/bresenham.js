function bresenham(x1, y1, x2, y2) {
  var m_new = 2 * (y2 - y1);
  var slope_error_new = m_new - (x2 - x1);

  for (x = x1, y = y1; x <= x2; x++) {
    console.log({ x, y });
    console.log({ slope_error_new, m_new });

    // Add slope to increment angle formed
    slope_error_new += m_new;

    // Slope error reached limit, time to
    // increment y and update slope error.
    if (slope_error_new >= 0) {
      y++;
      slope_error_new -= 2 * (x2 - x1);
    }
  }
}

// Driver code

var x1 = 3,
  y1 = 2,
  x2 = 15,
  y2 = 5;
bresenham(x1, y1, x2, y2);
