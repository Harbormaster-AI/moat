import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TrainingCourseService } from '../../../services/TrainingCourse.service';
import { TrainingCourse } from '../../../models/TrainingCourse';
import { SubBaseComponent } from '../../TrainingCourse/sub.base.component';

@Component({
    selector: 'app-create-trainingCourse',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTrainingCourseComponent extends SubBaseComponent implements OnInit {

    title = 'Add TrainingCourse';

    trainingCourseForm: FormGroup;
    trainingCourse: TrainingCourse;

    constructor( http: HttpClient,
        private trainingCourseService: TrainingCourseService,
        private fb: FormBuilder,
        private router: Router
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

    
    addTrainingCourse(code, title, durationHours, Prerequisites, Enrollments, JobProfiles, DeliveryMethod): void {
        this.trainingCourseService
        .addTrainingCourse(code, title, durationHours, Prerequisites, Enrollments, JobProfiles, DeliveryMethod)
            .subscribe(() => {
                this.router.navigate(['/indexTrainingCourse']);
            });
    }

    ngOnInit(): void {
    }
}