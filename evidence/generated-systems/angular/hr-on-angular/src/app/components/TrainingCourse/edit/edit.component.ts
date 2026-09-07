import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TrainingCourseService } from '../../../services/TrainingCourse.service';
import { SubBaseComponent } from '../../TrainingCourse/sub.base.component';


@Component({
    selector: 'app-edit-trainingCourse',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTrainingCourseComponent extends SubBaseComponent implements OnInit {

    title = 'Edit TrainingCourse';

    trainingCourseForm: FormGroup;
    trainingCourse: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TrainingCourseService,
        private fb: FormBuilder
) {
        super(http);
        this.trainingCourseForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  code: ['', Validators.required],
      title: ['', Validators.required],
      durationHours: ['', Validators.required],
      Prerequisites: ['', ],
      Enrollments: ['', ],
      JobProfiles: ['', ],
      DeliveryMethod: ['', ]
        });
    }

    
    updateTrainingCourse(code, title, durationHours, Prerequisites, Enrollments, JobProfiles, DeliveryMethod): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTrainingCourse(code, title, durationHours, Prerequisites, Enrollments, JobProfiles, DeliveryMethod, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTrainingCourse']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTrainingCourse(params['id']).subscribe(res => {
                this.trainingCourse = res;
            });
        });
    }
}