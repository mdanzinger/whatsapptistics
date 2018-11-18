var gulp = require('gulp'),
    gutil = require('gulp-util'),
    sass = require("gulp-sass"),
    autoprefixer = require('gulp-autoprefixer'),
    minifycss = require('gulp-minify-css'),
    uglify = require('gulp-uglify'),
    rename = require('gulp-rename'),
    clean = require('gulp-clean'),
    concat = require('gulp-concat'),
    notify = require('gulp-job');

// Copy files
gulp.task('copy', function() {
    return gulp.src('js/vendor/**/*.js')
        .pipe(gulp.dest('../dist/js/vendor'))
        .pipe(notify({ message: 'Copy task complete' }));
});

// Styles
gulp.task('styles', function() {
    return gulp.src('scss/**/*.scss')
        .pipe(sass().on('error', sass.logError))
        .pipe(autoprefixer('last 2 versions', '> 1%', 'Explorer 9', 'Firefox ESR', 'Opera 12.1'))
        .pipe(gulp.dest('../dist/css'))
        .pipe(rename(function (dir, base, ext) {
            return base + ".min" + ext;
        }))
        .pipe(minifycss())
        .pipe(gulp.dest('../dist/css'))
        .pipe(notify({ message: 'Styles task complete' }));
});

// Scripts
gulp.task('scripts', function() {
    return gulp.src(['js/*.js','!src/js/vendor/**/*.js'])
        .pipe(concat('main.js'))
        .pipe(gulp.dest('../dist/js'))
        .pipe(rename(function (dir, base, ext) {
            return base + ".min" + ext;
        }))
        .pipe(uglify())
        .pipe(gulp.dest('../dist/js'))
        .pipe(notify({ message: 'Scripts task complete' }));
});

// Clean
gulp.task('clean', function() {
    return gulp.src(['../dist/css', 'dist/js'], { read: false })
        .pipe(clean({force: true}));
});

// Default task
gulp.task('default', ['clean'], function() {
    gulp.start('copy', 'styles', 'scripts');
});

// Watch
gulp.task('watch', function() {

    // Watch .scss files
    gulp.watch('scss/**/*.scss', ['styles']);

    // Watch .js files
    gulp.watch('js/**/*.js', ['scripts']);

});