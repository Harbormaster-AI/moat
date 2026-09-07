import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TrainingEnrollmentService } from '../../../services/TrainingEnrollment.service';
import { TrainingEnrollment } from '../../../models/TrainingEnrollment';
import { SubBaseComponent } from '../../TrainingEnrollment/sub.base.component';

@Component({
    selector: 'app-create-trainingEnrollment',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTrainingEnrollmentComponent extends SubBaseComponent implements OnInit {

    title = 'Add TrainingEnrollment';

    trainingEnrollmentForm: FormGroup;
    trainingEnrollment: TrainingEnrollment;

    constructor( http: HttpClient,
        private trainingEnrollmentService: TrainingEnrollmentService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.trainingEnrollmentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  enrollmentNumber: ['', Validators.required],
      completionDate: ['', Validators.required],
      score: ['', Validators.required],
      Course: ['', ],
      Employee: ['', ],
      Instructor: ['', ],
      Status: ['', ]
        });
    }

    
    addTrainingEnrollment(enrollmentNumber, completionDate, score, Course, Employee, Instructor, Status): void {
        this.trainingEnrollmentService
        .addTrainingEnrollment(enrollmentNumber, completionDate, score, Course, Employee, Instructor, Status)
            .subscribe(() => {
                this.router.navigate(['/indexTrainingEnrollment']);
            });
    }

    ngOnInit(): void {
    }
}