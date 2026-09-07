
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TrainingCourseService } from '../../../services/TrainingCourse.service';
import { TrainingCourse } from '../../../models/TrainingCourse';

@Component({
    selector: 'app-index-trainingCourse',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTrainingCourseComponent implements OnInit {

    trainingCourses: TrainingCourse[] = [];

    constructor(
        private router: Router,
        private service: TrainingCourseService
) {}

    ngOnInit(): void {
        this.getTrainingCourses();
}

    getTrainingCourses(): void {
        this.service.getTrainingCourses().subscribe((res) => {
        this.trainingCourses = res;
    });
}

    deleteTrainingCourse(id: any): void {
        this.service.deleteTrainingCourse(id)
            .subscribe(() => {
                this.getTrainingCourses();
            });
    }
}