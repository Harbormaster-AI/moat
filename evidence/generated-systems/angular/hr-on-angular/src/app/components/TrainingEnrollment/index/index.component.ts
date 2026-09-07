
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TrainingEnrollmentService } from '../../../services/TrainingEnrollment.service';
import { TrainingEnrollment } from '../../../models/TrainingEnrollment';

@Component({
    selector: 'app-index-trainingEnrollment',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTrainingEnrollmentComponent implements OnInit {

    trainingEnrollments: TrainingEnrollment[] = [];

    constructor(
        private router: Router,
        private service: TrainingEnrollmentService
) {}

    ngOnInit(): void {
        this.getTrainingEnrollments();
}

    getTrainingEnrollments(): void {
        this.service.getTrainingEnrollments().subscribe((res) => {
        this.trainingEnrollments = res;
    });
}

    deleteTrainingEnrollment(id: any): void {
        this.service.deleteTrainingEnrollment(id)
            .subscribe(() => {
                this.getTrainingEnrollments();
            });
    }
}