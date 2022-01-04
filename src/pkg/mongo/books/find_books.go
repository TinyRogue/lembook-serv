package books

import (
	"context"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
)

var (
	authors      = []string{"Stanisław Lem", "Jerome Salinger", "Jacek Piekara", "Olga Tokarczuk", "Stefan Żeromski"}
	titles       = []string{"Solaris", "Buszujący W Zbożu", "Ja, Inkwizytor. Sługa boży. Cykl o Mordimerze Madderdinie. Tom 1", "Zgubiona dusza", "Dzieje grzechu"}
	descriptions = []string{
		"Kris Kelvin jest osobą, która ma wspomóc badania na obcej planecie, jednak od czasu przybycia na stację kosmiczną doświadcza dziwnych zjawisk. Ekipa przebywająca w bazie również zachowuje się podejrzanie, co jeszcze bardziej upewnia Kelvina, że dzieją się niepokojące rzeczy. Bohater podejrzewa, że w badaniach towarzyszy im obcy byt, którego ciężko zidentyfikować.",
		"Szesnastoletni Holden Caulfield, który nie może pogodzić się z otaczającą go głupotą, podłością, a przede wszystkim zakłamaniem, ucieka z college’u i przez kilka dni „buszuje” po Nowym Jorku, nim wreszcie powróci do domu rodzinnego. Historia tych paru dni, opowiedziana barwnym językiem, jest na pierwszy rzut oka zabawna, jednakże wkrótce spostrzegamy, że pod pozorami komizmu ważą się tutaj sprawy bynajmniej nie błahe...",
		"Minęło tysiąc pięćset lat, od kiedy Jezus zszedł z krzyża, utopił we krwi Jerozolimę i zdobył Rzym. Światem żądzą inkwizytorzy. Oto czas, w którym czarnoksiężnicy odprawiają bluźniercze rytuały, demony zstępują na świat, czarownice knują mroczne intrygi. A temu wszystkiemu musi przeciwstawić się Mordimer Madderdin - człowiek, którego serce jest tak gorące jak ogień stosów, na które posyła swe ofiary.",
		"Oto książka, która swoim głębokim przesłaniem zachwyci serce niejednego Czytelnika. \"Zgubiona dusza\" to historia o poszukiwaniu i cierpliwości, która jest z nim związana.\n\nNie jest trudno zgubić się w trudach dzisiejszego świata i zatracić we wszystkim, co nam proponuje. Przychodzi jednak w życiu moment na zastanowienie się, przebłysk, w trakcie którego możemy ponownie zweryfikować swoje wartości i fundamenty, na których planujemy odległą przyszłość. Czy na pewno są wystarczająco trwałe i dobre? Olga Tokarczuk w swojej najnowszej książce \"Zgubiona dusza\" porusza wiele istotnych problemów egzystencjalnych, z którymi stara się uporać.",
		"Powieść stała się skandalem literackim swojej epoki.Krytyka potępiła Żeromskiego za odważne ukazanie tematyki seksualnej i destrukcyjnej siły miłości. Wśród publiczności literackiej książka, nazywana „powieścią seksualną” lub „kolejowym romansem”, stała się niezwykle popularna, doczekała się także kilku ekranizacji.\nTo historia tragicznej miłości, opowieść o namiętności, obsesji, zdradzie i zbrodni.",
	}
	coverURLs = []string{
		"https://bigimg.taniaksiazka.pl/images/popups/8DD/9788308070512.jpg",
		"https://bigimg.taniaksiazka.pl/images/popups/E0A/9788381255004.jpg",
		"https://bigimg.taniaksiazka.pl/images/popups/36F/9788379640058.jpg",
		"https://bigimg.taniaksiazka.pl/images/popups/66F/9788361488743.jpg",
		"https://bigimg.taniaksiazka.pl/images/popups/D71/@9788362121533_31.jpg",
	}
	genres = []string{"Science-Fiction", "Powieść", "Fantasy", "Polska", "Polska"}
)

func (s *Service) FindBooks(ctx context.Context, b *model.WhatBook) (model.Books, error) {
	var templateBooks model.Books
	for i, author := range authors {
		templateBooks.Books = append(templateBooks.Books, &model.Book{
			Author:      &author,
			Title:       &titles[i],
			Description: &descriptions[i],
			CoverURL:    &coverURLs[i],
			Genre:       &genres[i],
		})
	}
	return templateBooks, nil
}
