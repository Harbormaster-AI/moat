import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TrainingEnrollmentService } from '../../../services/TrainingEnrollment.service';
import { SubBaseComponent } from '../../TrainingEnrollment/sub.base.component';


@Component({
    selector: 'app-edit-trainingEnrollment',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTrainingEnrollmentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit TrainingEnrollment';

    trainingEnrollmentForm: FormGroup;
    trainingEnrollment: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TrainingEnrollmentService,
        private fb: FormBuilder
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

    
    updateTrainingEnrollment(enrollmentNumber, completionDate, score, Course, Employee, Instructor, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTrainingEnrollment(enrollmentNumber, completionDate, score, Course, Employee, Instructor, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTrainingEnrollment']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTrainingEnrollment(params['id']).subscribe(res => {
                this.trainingEnrollment = res;
            });
        });
    }
}