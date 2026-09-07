import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InterviewService } from '../../../services/Interview.service';
import { SubBaseComponent } from '../../Interview/sub.base.component';


@Component({
    selector: 'app-edit-interview',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInterviewComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Interview';

    interviewForm: FormGroup;
    interview: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InterviewService,
        private fb: FormBuilder
) {
        super(http);
        this.interviewForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  interviewDate: ['', Validators.required],
      feedback: ['', Validators.required],
      Requisition: ['', ],
      Candidate: ['', ],
      Interviewers: ['', ],
      Stage: ['', ],
      Result: ['', ]
        });
    }

    
    updateInterview(interviewDate, feedback, Requisition, Candidate, Interviewers, Stage, Result): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInterview(interviewDate, feedback, Requisition, Candidate, Interviewers, Stage, Result, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInterview']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInterview(params['id']).subscribe(res => {
                this.interview = res;
            });
        });
    }
}