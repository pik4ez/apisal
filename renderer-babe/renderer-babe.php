#!/usr/bin/env php
<?php

$speed = 6; // sec


$in = file_get_contents('php://stdin');
$objects = array_map(function ($str) {
    return json_decode($str, true);
}, explode("\n", $in));

//print_r($objects);
$spans = [];
$divs = [];
$list = [];

$index = 0;
foreach ($objects as $object) {
    if (!$object['images'] || !$object['title']) {
        continue;
    }
//    continue;

    $title = htmlspecialchars($object['title']);

    $imagesCount = 0;
    foreach ($object['images'] as $image) {
        $sec = $index * $speed;
        $childIndex = $index + 1;

        $url = $image['url'];
//        $url = strtr($url, ['http://' => '//', 'https://' => '//']);
        $url = htmlspecialchars($url);

        $spans[] = <<<CSS
.cb-slideshow li:nth-child({$childIndex}) span {
    background-image: url("{$url}"); -webkit-animation-delay: {$sec}s; -moz-animation-delay: {$sec}s;
    -o-animation-delay: {$sec}s; -ms-animation-delay: {$sec}s; animation-delay: {$sec}s;
}
CSS;
        if ($childIndex > 1) {
            $divs[] = <<<CSS
.cb-slideshow li:nth-child({$childIndex}) div {
    -webkit-animation-delay: {$sec}s; -moz-animation-delay: {$sec}s;
    -o-animation-delay: {$sec}s; -ms-animation-delay: {$sec}s;
    animation-delay: {$sec}s;
}
CSS;
        }

        $list[] = <<<HTML
<li><span></span><div><h3>{$title}</h3></div></li>
HTML;


        $index++;
        $imagesCount++;

        if ($index >= 6) {
            break 2;
        }
        
        if ($imagesCount >= 1) {
            break 1;
        }
    }

}

?>
<!DOCTYPE html>
<!--[if lt IE 7 ]>
<html class="ie ie6 no-js" lang="en"> <![endif]-->
<!--[if IE 7 ]>
<html class="ie ie7 no-js" lang="en"> <![endif]-->
<!--[if IE 8 ]>
<html class="ie ie8 no-js" lang="en"> <![endif]-->
<!--[if IE 9 ]>
<html class="ie ie9 no-js" lang="en"> <![endif]-->
<!--[if gt IE 9]><!-->
<html class="no-js" lang="en"><!--<![endif]-->
<head>
    <meta charset="UTF-8"/>
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <title>АПИ Сал: movie</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description"
          content="Fullscreen Background Image Slideshow with CSS3 - A Css-only fullscreen background image slideshow"/>
    <meta name="keywords" content="css3, css-only, fullscreen, background, slideshow, images, content"/>
    <meta name="author" content="Codrops"/>
    <link rel="stylesheet" type="text/css" href="./static/babe/css/demo.css"/>
    <script type="text/javascript" src="./static/babe/js/modernizr.custom.86080.js"></script>
    <style>
        .cb-slideshow,
        .cb-slideshow:after {
            position: fixed;
            width: 100%;
            height: 100%;
            top: 0px;
            left: 0px;
            z-index: 0;
        }

        .cb-slideshow:after {
            content: '';
            background: transparent url(./static/babe/images/pattern.png) repeat top left;
        }

        .cb-slideshow li span {
            width: 100%;
            height: 100%;
            position: absolute;
            top: 0px;
            left: 0px;
            color: transparent;
            background-size: cover;
            background-position: 50% 50%;
            background-repeat: none;
            opacity: 0;
            z-index: 0;
            -webkit-backface-visibility: hidden;
            -webkit-animation: imageAnimation <?php print $sec + $speed; ?>s linear infinite 0s;
            -moz-animation: imageAnimation <?php print $sec + $speed; ?>s linear infinite 0s;
            -o-animation: imageAnimation <?php print $sec + $speed; ?>s linear infinite 0s;
            -ms-animation: imageAnimation <?php print $sec + $speed; ?>s linear infinite 0s;
            animation: imageAnimation <?php print $sec + $speed; ?>s linear infinite 0s;
        }

        .cb-slideshow li div {
            z-index: 1000;
            position: absolute;
            bottom: 30px;
            left: 0px;
            width: 100%;
            text-align: center;
            opacity: 0;
            color: #fff;
            -webkit-animation: titleAnimation <?php print $sec + $speed; ?>s linear infinite 0s;
            -moz-animation: titleAnimation <?php print $sec + $speed; ?>s linear infinite 0s;
            -o-animation: titleAnimation <?php print $sec + $speed; ?>s linear infinite 0s;
            -ms-animation: titleAnimation <?php print $sec + $speed; ?>s linear infinite 0s;
            animation: titleAnimation <?php print $sec + $speed; ?>s linear infinite 0s;
        }

        .cb-slideshow li div h3 {
            font-family: 'Arial Narrow', Arial, sans-serif;
            font-size: 120px;
            padding: 0;
            line-height: 200px;
        }

        <?php
            print implode("\n", $spans);
            print implode("\n", $divs);
        ?>
        /* Animation for the slideshow images */
        @-webkit-keyframes imageAnimation {
            0% {
                opacity: 0;
                -webkit-animation-timing-function: ease-in;
            }
            8% {
                opacity: 1;
                -webkit-animation-timing-function: ease-out;
            }
            17% {
                opacity: 1
            }
            25% {
                opacity: 0
            }
            100% {
                opacity: 0
            }
        }

        @-moz-keyframes imageAnimation {
            0% {
                opacity: 0;
                -moz-animation-timing-function: ease-in;
            }
            8% {
                opacity: 1;
                -moz-animation-timing-function: ease-out;
            }
            17% {
                opacity: 1
            }
            25% {
                opacity: 0
            }
            100% {
                opacity: 0
            }
        }

        @-o-keyframes imageAnimation {
            0% {
                opacity: 0;
                -o-animation-timing-function: ease-in;
            }
            8% {
                opacity: 1;
                -o-animation-timing-function: ease-out;
            }
            17% {
                opacity: 1
            }
            25% {
                opacity: 0
            }
            100% {
                opacity: 0
            }
        }

        @-ms-keyframes imageAnimation {
            0% {
                opacity: 0;
                -ms-animation-timing-function: ease-in;
            }
            8% {
                opacity: 1;
                -ms-animation-timing-function: ease-out;
            }
            17% {
                opacity: 1
            }
            25% {
                opacity: 0
            }
            100% {
                opacity: 0
            }
        }

        @keyframes imageAnimation {
            0% {
                opacity: 0;
                animation-timing-function: ease-in;
            }
            8% {
                opacity: 1;
                animation-timing-function: ease-out;
            }
            17% {
                opacity: 1
            }
            25% {
                opacity: 0
            }
            100% {
                opacity: 0
            }
        }

        /* Animation for the title */
        @-webkit-keyframes titleAnimation {
            0% {
                opacity: 0
            }
            8% {
                opacity: 1
            }
            17% {
                opacity: 1
            }
            19% {
                opacity: 0
            }
            100% {
                opacity: 0
            }
        }

        @-moz-keyframes titleAnimation {
            0% {
                opacity: 0
            }
            8% {
                opacity: 1
            }
            17% {
                opacity: 1
            }
            19% {
                opacity: 0
            }
            100% {
                opacity: 0
            }
        }

        @-o-keyframes titleAnimation {
            0% {
                opacity: 0
            }
            8% {
                opacity: 1
            }
            17% {
                opacity: 1
            }
            19% {
                opacity: 0
            }
            100% {
                opacity: 0
            }
        }

        @-ms-keyframes titleAnimation {
            0% {
                opacity: 0
            }
            8% {
                opacity: 1
            }
            17% {
                opacity: 1
            }
            19% {
                opacity: 0
            }
            100% {
                opacity: 0
            }
        }

        @keyframes titleAnimation {
            0% {
                opacity: 0
            }
            8% {
                opacity: 1
            }
            17% {
                opacity: 1
            }
            19% {
                opacity: 0
            }
            100% {
                opacity: 0
            }
        }

        /* Show at least something when animations not supported */
        .no-cssanimations .cb-slideshow li span {
            opacity: 1;
        }

        @media screen and (max-width: 1140px) {
            .cb-slideshow li div h3 {
                font-size: 100px
            }
        }

        @media screen and (max-width: 600px) {
            .cb-slideshow li div h3 {
                font-size: 60px
            }
        }        </style>
</head>
<body id="page">
<ul class="cb-slideshow">
<?php
    print implode("\n", $list);
?>
</ul>
</body>
</html>

