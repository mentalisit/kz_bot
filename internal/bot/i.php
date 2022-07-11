<?php
// топ за 7 дней

if (preg_match("/^(Rs|rs)\s(Top|top)\s(W|w)$/i", $mtext, $m)){

	$time=time();
	$ttime = $time - 604800;
	$number = 1;
	$mesage="\xF0\x9F\x93\x96 ТОП15 Участников за 7 дней:";

	$result = $mysqli->query("SELECT name FROM sborkz WHERE time>'$ttime' AND chatid='$cid' AND active='1' GROUP BY name");
		if ($result->num_rows > 0) {
             while($row = $result->fetch_assoc()) {
			  $result2 = $mysqli->query("SELECT * FROM sborkz WHERE name='$row[name]' AND time>'$ttime' AND chatid='$cid' AND active='1'");
	          $result3 = $mysqli->query("INSERT INTO temptop (name,numkz) VALUES ('$row[name]','$result2->num_rows')");
			 }

     $result = $mysqli->query("SELECT * FROM temptop ORDER BY numkz DESC LIMIT 15");
	    while($row = $result->fetch_assoc()) {
          $mesage2 = $mesage2."$number. $row[name] - ($row[numkz])\n";
		  $number = $number+1;
           }
          $bot->sendMessage($message->getChat()->getId(), "$mesage\n$mesage2");
     $result2 = $mysqli->query("DELETE FROM temptop");
      $number = 1;

		} else {

$bot->sendMessage($message->getChat()->getId(), "ТОП-лист пуст");

		}




}
$mysqli->close();
}

//

// топ за сутки

if (preg_match("/^(Rs|rs)\s(Top|top)\s(D|d)$/i", $mtext, $m)){

                   if($nameid == "") {
				      $bot->sendMessage($message->getChat()->getId(), "Для того что бы БОТ мог Вас индентифицировать, создайте уникальный НикНей в настройках. Вы можете использовать a-z, 0-9 и символы подчеркивания. Минимальная длина - 5 символов.");
	               } else {


// вносим в базу значение просмотров ТОПа конкретным пользователем, если есть уже увеличиваем на 1
$resulttoptop = $mysqli->query("SELECT name FROM toptoptop WHERE name='$nameid' AND cid='$cid'");
if ($resulttoptop->num_rows > 0 ) {

	 $resulttop = $mysqli->query("UPDATE toptoptop SET count=count + 1 WHERE name='$nameid' AND cid='$cid'");


} else {

      $resulttop = $mysqli->query("INSERT INTO toptoptop (name,cid,count) VALUES ('$nameid','$cid','1')");

}
//Конец работы с ТОП




	$time=time();
	// вычисляем время от полуночи сегодняшнего дня
	$ttime = (new DateTime('today midnight'))->getTimestamp();
	$number = 1;
	$mesage="\xF0\x9F\x93\x96 ТОП10 Участников за текущие сутки:";


	$result = $mysqli->query("SELECT name FROM sborkz WHERE time>'$ttime' AND chatid='$cid' AND active='1' GROUP BY name");
		if ($result->num_rows > 0) {
             while($row = $result->fetch_assoc()) {
			  $result2 = $mysqli->query("SELECT * FROM sborkz WHERE name='$row[name]' AND time>'$ttime' AND chatid='$cid' AND active='1'");
	          $result3 = $mysqli->query("INSERT INTO temptop (name,numkz) VALUES ('$row[name]','$result2->num_rows')");
			 }

     $result = $mysqli->query("SELECT * FROM temptop ORDER BY numkz DESC LIMIT 10");
	    while($row = $result->fetch_assoc()) {
          $mesage2 = $mesage2."$number. $row[name] - ($row[numkz])\n";
		  $number = $number+1;
           }
          $bot->sendMessage($message->getChat()->getId(), "$mesage\n$mesage2");
     $result2 = $mysqli->query("DELETE FROM temptop");
      $number = 1;

		} else {

$bot->sendMessage($message->getChat()->getId(), "ТОП-лист пуст");

		}




}
$mysqli->close();
}

// начало блока rs sub n
if (preg_match("/^(rs|Rs)\s(Sub|sub)\s([0-9]+)$/i", $mtext, $m)){

$idkz = $m[3]; //номер КЗ

//ищем активную кз с указанным id и состоит ли пользователь в ней
$result = $mysqli->query("SELECT lvlkz FROM sborkz WHERE chatid='$cid' AND numberkz='$idkz' AND active='1' AND name='$nameid' ORDER BY id DESC LIMIT 1");
$lvlkz = $result->fetch_assoc();
$lvlkz = $lvlkz[lvlkz];

	  if ($result->num_rows > 0) {


	$resultsubs = $mysqli->query("SELECT name FROM subscribe WHERE lvlkz='$lvlkz' AND chatid='$cid' AND name!='$nameid'");
               $count = $resultsubs->num_rows;
               $subs = '';
			   $i = 0;
             while ($i<=$count) {

			 $resultsubss = $mysqli->query("SELECT name FROM subscribe WHERE lvlkz='$lvlkz' AND chatid='$cid' AND name!='$nameid' LIMIT $i,5");

			   if ($resultsubss->num_rows > 0 ) {

				  while($row = $resultsubss->fetch_assoc()) {

				    $subs = $subs." @".$row[name].",";

               }

			    }
			//if ($subs = '') {} else { $subs = substr_replace($subs,'.',-1);}
			 $bot->sendMessage($message->getChat()->getId(), "$subs");
			 $subs = '';
		    	$i=$i+5;
			}
	  $bot->sendMessage($message->getChat()->getId(), "Друзья! На КЗ $lvlkz игрок не явился в игру, у кого есть желание заменить его?");

	  } else {
$bot->sendMessage($message->getChat()->getId(), "КЗ с номером $idkz не найдена, либо Вы не являетесь ее участником.", "html", true);
	  }
$mysqli->close();
	}
// конец блока rs sub n


// блок замены игрока в очеред n$mesage
if (preg_match("/^(rs|Rs)\s(Sub|sub)\s([0-9]+)\s(\@\w+)\s(\@\w+)$/i", $mtext, $m)){

$idkz = $m[3]; //номер КЗ
$pdel = substr($m[4],1); //игрок на удаление
$pnew = substr($m[5],1); //новый игрок
// если ивент неактивен переменная имеет значение numberkz и поиск ведется по этому полю
// если ивент активен, переменная меняет значение на idkz , и поиск ведется по этому полю
$polekz = "numberkz";
$formmessage3 = "Номер КЗ - $idkz.";

$resultevent = $mysqli->query("SELECT numevent FROM rsevent WHERE chatid='$cid' AND activeevent='1' ORDER BY numevent DESC LIMIT 1");
	    if ($resultevent->num_rows > 0) {
		$numberevent = $resultevent->fetch_assoc();
		$numberevent =  $numberevent[numevent];
		$polekz = "idkz";
		$formmessage3 = "ID КЗ - $idkz.";
		}




// ищем номер КЗ с состоянием закрыта и наличием человека в этой очереди
      $result = $mysqli->query("SELECT name,lvlkz FROM sborkz WHERE chatid='$cid' AND $polekz='$idkz' AND active='1' AND name='$nameid' ORDER BY id DESC LIMIT 1");
        if ($result->num_rows > 0) {
			 $lvlkz = $result->fetch_assoc();
			 $lvlkz = $lvlkz[lvlkz];
           // ищем игрока на удаление из этой очеред
           $result2 = $mysqli->query("SELECT id FROM sborkz WHERE chatid='$cid' AND $polekz='$idkz' AND active='1' AND name='$pdel' ORDER BY id DESC LIMIT 1");
           $idname = $result2->fetch_assoc();
		   $idname =  $idname[id];
		   if ($result2->num_rows > 0) {

				$result3 = $mysqli->query("UPDATE sborkz SET name='$pnew',nameid='0' WHERE id='$idname' AND $polekz='$idkz' AND chatid='$cid' AND lvlkz='$lvlkz'");
                $bot->sendMessage($message->getChat()->getId(), "Игрок под ником $pnew успешно внесен в базу данных.", "html", true);


// вывод списка очереди под номером n
		        $result5 = $mysqli->query("SELECT * FROM sborkz WHERE $polekz='$idkz' AND chatid='$cid' AND active='1' AND numberevent='$numberevent' AND lvlkz='$lvlkz'");

				if ($result5->num_rows > 0) {
                  while($row = $result5->fetch_assoc()) {


// выковыриваем из базы значение количества походов на кз
                $result6 = $mysqli->query("SELECT name FROM sborkz WHERE name='$row[name]' AND chatid='$cid' AND lvlkz='$row[lvlkz]' AND active='1'");
                  if ($result6->num_rows < 1) {
                          $numberkzz = 0;
				  } else {
				          $numberkzz = $result6->num_rows;

			   }
// конец проверки колва походов
// выковыриваем из базы иконки
     $resulticon = $mysqli->query("SELECT icon1,icon2,icon3,icon4 FROM userinfo WHERE name='$row[name]' AND chatid='$cid'");
     $usericon = $resulticon->fetch_assoc();
	 $icon1 = $usericon[icon1];
	 $icon2 = $usericon[icon2];
	 $icon3 = $usericon[icon3];
	 $icon4 = $usericon[icon4];
// конец иконкам )))


                $formmessage1 = "Обновленный состав на КЗ $row[lvlkz]:\n";

				$proba=$proba+1;
				$formmessage2 = $formmessage2."$proba. <b>$row[name]</b> ($numberkzz) $icon1$icon2$icon3$icon4\n";


							  }

				$bot->sendMessage($message->getChat()->getId(), "$formmessage1$formmessage2$formmessage3", "html", true);
                $proba=0;
				$formmessage2="";

				} else {
					$countkz = $countkz-1;
			           }
// конец вывода очереди после изменения игрока


			 } else {
                $bot->sendMessage($message->getChat()->getId(), "Игрок под ником $pdel в данной очереди не найден.", "html", true);
             }




        } else {
		$bot->sendMessage($message->getChat()->getId(), "Очередь под номер $idkz не найдена, либо Вы не являетесь ее участником.", "html", true);
		}

}
// конец блока


// блок удаления игрока из очереди
if (preg_match("/^(rs|Rs)\s(Kick|kick)\s(\@\w+)$/i", $mtext, $m)){
 if($nameid == "") {
				      $bot->sendMessage($message->getChat()->getId(), "Для того что бы БОТ мог Вас индентифицировать, создайте уникальный НикНей в настройках. Вы можете использовать a-z, 0-9 и символы подчеркивания. Минимальная длина - 5 символов.");
	               } else {






$nicknamedel = substr($m[3],1); //игрок на удаление



	if ($nicknamedel == $nameid) {
$bot->sendMessage($message->getChat()->getId(), "Просто выйди из очереди, Злодей", "html", true);
	} else {

// ищем номер КЗ с состоянием открыта и наличием человека в этой очереди
      $result = $mysqli->query("SELECT * FROM sborkz WHERE chatid='$cid' AND active='0' AND name='$nameid'");
        if ($result->num_rows > 0) {

		  while($rows = $result->fetch_assoc()) {
			$lvlkz = $rows[lvlkz];


    // ищем игрока на удаление из этой очеред
           $result2 = $mysqli->query("SELECT id FROM sborkz WHERE chatid='$cid' AND lvlkz='$lvlkz' AND active='0' AND name='$nicknamedel'");
           $idname = $result2->fetch_assoc();
		   $idname = $idname[id];

		   if ($result2->num_rows > 0) {

				$result3 = $mysqli->query("DELETE FROM sborkz WHERE id='$idname' AND chatid='$cid' AND lvlkz='$lvlkz' AND active='0'");
                $bot->sendMessage($message->getChat()->getId(), "Игрок под ником <b>$nicknamedel</b> успешно удален из очереди на КЗ $lvlkz уровня.", "html", true);





// вывод списка очереди под номером n
		        $result5 = $mysqli->query("SELECT * FROM sborkz WHERE lvlkz='$lvlkz' AND chatid='$cid' AND active='0'");

				if ($result5->num_rows > 0) {
                  while($row = $result5->fetch_assoc()) {


// выковыриваем из базы значение количества походов на кз
                $result6 = $mysqli->query("SELECT name FROM sborkz WHERE name='$row[name]' AND chatid='$cid' AND lvlkz='$row[lvlkz]' AND active='1'");
                  if ($result6->num_rows < 1) {
                          $numberkzz = 0;
				  } else {
				          $numberkzz = $result6->num_rows;

			   }
// конец проверки колва походов
// выковыриваем из базы иконки
     $resulticon = $mysqli->query("SELECT icon1,icon2,icon3,icon4 FROM userinfo WHERE name='$row[name]' AND chatid='$cid'");
     $usericon = $resulticon->fetch_assoc();
	 $icon1 = $usericon[icon1];
	 $icon2 = $usericon[icon2];
	 $icon3 = $usericon[icon3];
	 $icon4 = $usericon[icon4];
// конец иконкам )))


                $formmessage1 = "Очередь КЗ $row[lvlkz]:\n";

				$proba=$proba+1;
				$formmessage2 = $formmessage2."$proba. <b>$row[name]</b> ($numberkzz) $icon1$icon2$icon3$icon4\n";


							  }

				$bot->sendMessage($message->getChat()->getId(), "$formmessage1$formmessage2$formmessage3", "html", true);
                $proba=0;
				$formmessage2="";

				}


			}

}
} else {
		$bot->sendMessage($message->getChat()->getId(), "Вы не являетесь участником очереди на КЗ.", "html", true);
		}

  }
 }
 $mysqli->close();
}
// конец блока



if (preg_match("/^(About|about)\s(bot|Bot)$/i", $mtext, $m)){

$bot->sendMessage($message->getChat()->getId(), "БОТ для сбора на КЗ. Корпорация СОЮЗ (2021 год).\n\nИдейный
вдохновитель - <b>Error_256</b>.\nРазработка и поддержка @ArtZor_hs", "html", true);

}


if (preg_match("/^(rs|Rs)\s(Event|event)$/i", $mtext, $m)){
	if($nameid == "") {
 $bot->sendMessage($message->getChat()->getId(), "Для того что бы БОТ мог Вас индентифицировать, создайте уникальный НикНей в настройках. Вы можете использовать a-z, 0-9 и символы подчеркивания. Минимальная длина - 5 символов.");
 } else {




	$result = $mysqli->query("SELECT * FROM rsevent WHERE chatid='$cid' AND activeevent='0'");
      if ($result->num_rows > 0) {
          $numevent = $result->num_rows;
		  $numberevent = 1;

		  $message1 = 'Ваша корпорация провела '.$numevent.' событи(е,й,я) для частных красных звезд ( Ивент ).';
		  $message2 = '';

		  $resultevent = $mysqli->query("SELECT * FROM rsevent WHERE chatid='$cid' AND activeevent='0'");
		  while($row = $resultevent->fetch_assoc()) {
		  $datestartnorm = date("d.m.Y", strtotime($row[datestart]));
		  $datestopnorm = date("d.m.Y", strtotime($row[datestop]));
		  $resultnumkz = $mysqli->query("SELECT numberevent FROM sborkz WHERE chatid='$cid' AND numberevent='$row[numevent]' GROUP BY idkz");
		  $resultnumkzz = $resultnumkz->num_rows;
		  $message2 = $message2."\n№$row[numevent] проходил c $datestartnorm по $datestopnorm / $resultnumkzz КЗ";

			$numberevent = $numberevent + 1;
			}

		  $message3 = "\nТОП по каждому Ивенту доступен по команде <b>rs top event n</b> , где n - номер Ивента.";
		  $bot->sendMessage($message->getChat()->getId(), "$message1\n$message2\n$message3", "html", true);


	  } else {

$bot->sendMessage($message->getChat()->getId(), "Информация в базе данных не найдена ", "html", true);

	}
 }
 $mysqli->close();
}


// массовая отправка сообщений во все чаты
//--------------------------------------------------
if (preg_match("/^(Message|message)\s(\"((?<=[\"'])[^\"']+)\")$/i", $mtext, $m)){


      if($nameid == "ArtZor_hs") {




   $messsage = $m[2];

	$result = $mysqli->query("SELECT chatid FROM sborkz GROUP BY chatid");



	 if ($result->num_rows > 0) {
       while($row = $result->fetch_assoc()) {

			$ccid = $row[chatid];




	   }

	$bot->sendMessage($message->getChat()->getId(), "Сообщение отправлено!!!");

} }
// конец блока
// -----------------------------------------------






//соло катка
//------------------------------------------------------------------
if (preg_match("/^(Rs|rs)\s(Solo|solo)\s([4-9]|[1][0-1])$/", $mtext, $m)){

		           if($nameid == "") {
				      $bot->sendMessage($message->getChat()->getId(), "Для того что бы БОТ мог Вас индентифицировать, создайте уникальный НикНей в настройках. Вы можете использовать a-z, 0-9 и символы подчеркивания. Минимальная длина - 5 символов.");
	               } else {
		//уровень кз
         $lvlkz=$m[3];
        // время добавления в базу
        $time=time();
        $date=date('Y-m-d');

$resultivent = $mysqli->query("SELECT numevent FROM rsevent WHERE chatid='$cid' AND activeevent='1'");
	    if ($resultivent->num_rows > 0) {




// перед добавлением в очередь считываем порядковый номер КЗ, если такового нет, добавляем:
$result = $mysqli->query("SELECT number FROM numkz WHERE lvlkz='$lvlkz' AND chatid='$cid'");
if ($result->num_rows < 1 ) {
	$result = $mysqli->query("INSERT INTO numkz (lvlkz,number,chatid) VALUES ('$lvlkz','1','$cid')");
} else {
$numberkz = $result->fetch_assoc();
$numberkz = $numberkz[number];
}
//Конец проверки наличия порядкового номера КЗ.


//считываем id катки, которая после будет использоваться для внесения очей в базу
//
//

$numevent = $resultivent->fetch_assoc();
$numevent = $numevent[numevent];

// перед добавлением в очередь считываем id катки, если таковой нет, добавляем:
$result = $mysqli->query("SELECT numberkz FROM idkzivent WHERE chatid='$cid'");
if ($result->num_rows < 1 ) {
	$result = $mysqli->query("INSERT INTO idkzivent (chatid,numberkz) VALUES ('$cid','1')");

} else {
$idkz = $result->fetch_assoc();
$idkz = $idkz[numberkz];
}
//Конец проверки наличия порядкового номера КЗ.


 if (!$mysqli->query("INSERT INTO sborkz (name, nameid, mesid, chatid, time, date, lvlkz, nofarm, numberkz, idkz, numberevent, eventpoints, active, timedown, activedel) VALUES ('$nameid', '$nameidid', '$mesid', '$cid', '$time', '$date', '$lvlkz', '2', '$numberkz', '$idkz', '$numevent', '0','1','30', '0')"))
            {
          $$bot->sendMessage($message->getChat()->getId(), "Проблема с доступом к Базе Данных. Сообщие ошибку @ArtZor_hs.");
            }



// выковыриваем из базы значение количества походов на кз
                $result3 = $mysqli->query("SELECT * FROM sborkz WHERE name='$nameid' AND chatid='$cid' AND lvlkz='$lvlkz' AND active='1'");
                  if ($result3->num_rows < 1) {
                          $numberkzz = 0;
				  } else {
				          $numberkzz = $result3->num_rows;

			   }
// конец проверки колва походов
// выковыриваем из базы значение количества походов на кз
                $result3 = $mysqli->query("SELECT * FROM sborkz WHERE name='$nameid' AND chatid='$cid' AND lvlkz='$lvlkz' AND active='1'");
                  if ($result3->num_rows < 1) {
                          $numberkzz = 0;
				  } else {
				          $numberkzz = $result3->num_rows;

			   }
// конец проверки колва походов




				$messsage = "\xE2\x9C\x85 Соло поход на КЗ $lvlkz:\n\xF0\x9F\x91\x89 @$nameid ($numberkzz)\n\xE2\x9D\x97ID КЗ - $idkz\n$resulttime\nУдачи!";


				$bot->sendMessage($message->getChat()->getId(), "$messsage", "html", true);


// вносим в базу порядковый номер кз
			   $result2 = $mysqli->query("UPDATE sborkz SET numberkz='$numberkz', idkz='$idkz', numberevent='$numevent', activedel='0' WHERE name='$nameid' AND lvlkz='$lvlkz' AND chatid='$cid' AND active='0'");
			   $result3 = $mysqli->query("DELETE FROM sborkz WHERE name='$nameid' AND lvlkz!='$lvlkz' AND chatid='$cid' AND active='0'");

               // ищем по имени игрока в других чатах, если есть выводим сообщение что он ушел на КЗ, и удаляем его

			   $resultsearchchat = $mysqli->query("SELECT * FROM sborkz WHERE name='$row[name]' AND chatid!='$cid' AND active='0'");
			    if ($resultsearchchat->num_rows > 0 ) {

				  while($row = $resultsearchchat->fetch_assoc()) {

				    $info = file_get_contents("https://api.telegram.org/bot1962860522:AAGZ0z5NHebzRLzeLe_69dYAnuzGuoaNPvM/sendmessage?chat_id=$row[chatid]&text=\xE2\x9D\x97@$row[name] ушел на КЗ в другом чате");


                    //удаление игрока из всех чередей в других чатах
			        $resultdelallchat = $mysqli->query("DELETE FROM sborkz WHERE name='$row[name]' AND chatid!='$cid' AND active='0'");

               }

			    }


			//	$result = $mysqli->query("UPDATE sborkz SET active='1', activedel='0' WHERE lvlkz='$lvlkz' AND chatid='$cid'");
				$result = $mysqli->query("UPDATE numkz SET number=number + 1 WHERE lvlkz='$lvlkz' AND chatid='$cid'");




		$result = $mysqli->query("UPDATE idkzivent SET numberkz=numberkz + 1 WHERE chatid='$cid'");

	// конец проверки

			}	 else {

$bot->sendMessage($message->getChat()->getId(), "Команда доступна во время активного События частных КЗ ( Ивента ).");


				   }			}

		 $mysqli->close();

       }

// конец блока
//------------------------------------------------------------------


?>