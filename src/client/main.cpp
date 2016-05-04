/*
 * This file is part of "The Scoreboard"
 * Copyright (C) 2016  Tobias Polzer

 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 * 
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * 
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
 */
#include <QGuiApplication>
#include <QQmlApplicationEngine>
#include <QtQml>
#include <QQuickWindow>
#include <QVariant>
#include <QMetaObject>
#include <QJsonArray>
#include <QTcpSocket>
#include <QtEndian>
#include "ScoreboardClient.h"
#include "clock.h"

int main(int argc, char *argv[])
{
    QGuiApplication app(argc, argv);

    QQmlApplicationEngine engine;

	qmlRegisterType<Clock>("de.bulsa.clock", 0, 1, "Clock");

    engine.load(QUrl(QStringLiteral("qrc:/main.qml")));

    QQuickWindow *obj = qobject_cast<QQuickWindow*>(engine.rootObjects().first());
	ScoreboardClient client(engine);
	client.run();

	QObject::connect(&client, SIGNAL(contestSetup(QVariant,QVariant,QVariant)),
			obj, SLOT(contestSetup(QVariant,QVariant,QVariant)));


    return app.exec();
}
