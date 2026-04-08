// Copyright (C) 2025 The Qt Company Ltd.
// SPDX-License-Identifier: MIT

#include <QtGui>

class GuiTypes : public QObject
{
    Q_OBJECT

public:
    using QObject::QObject;

    QColor qColorRgb = QColor(255, 0, 128);
    QColor qColorRgba = QColor(255, 0, 128, 100);
    QColor qColorHsv = QColor::fromHsv(210, 200, 180);
    QColor qColorCmyk = QColor::fromCmyk(100, 200, 50, 25);
    QColor qColorHsl = QColor::fromHsl(120, 200, 150);
    QColor qColorInvalid = QColor();
    QColor qColorNamed = QColor("steelblue");

    QBitmap qBitmap = QBitmap(":/qt.png");
    QBrush qBrush = QBrush();
    QCursor qCursor = QCursor();
    QFont qFont = QFont();
    QImage qImage = QImage(":/qt.png");
    QKeySequence qKeySequence = QKeySequence(Qt::CTRL + Qt::Key_Q);
    QMatrix4x4 qMatrix4x4 = QMatrix4x4();
    QPalette qPalette = QPalette();
    QPen qPen = QPen();
    QPicture qPicture = QPicture();
    QPixmap qPixmap = QPixmap(":/qt.png");
    QPolygon qPolygon = QPolygon({QPoint(0, 0), QPoint(42, 0), QPoint(42, 42), QPoint(0, 42)});
    QPolygonF qPolygonF = QPolygonF({QPointF(0.5, 0.5), QPointF(4.2, 0.5), QPointF(4.2, 4.2), QPointF(0.5, 4.2)});
    QQuaternion qQuaternion = QQuaternion(0., 0., 0., 1.);
    QRegion qRegion = QRegion();
    QTransform qTransform = QTransform();
    QVector2D qVector2D = QVector2D(42., 43.);
    QVector3D qVector3D = QVector3D(42., 43., 44.);
    QVector4D qVector4D = QVector4D(42., 43., 44., 45.);
};
