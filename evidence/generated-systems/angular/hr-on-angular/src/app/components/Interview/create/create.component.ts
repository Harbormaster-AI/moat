import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InterviewService } from '../../../services/Interview.service';
import { Interview } from '../../../models/Interview';
import { SubBaseComponent } from '../../Interview/sub.base.component';

@Component({
    selector: 'app-create-interview',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInterviewComponent extends SubBaseComponent implements OnInit {

    title = 'Add Interview';

    interviewForm: FormGroup;
    interview: Interview;

    constructor( http: HttpClient,
        private interviewService: InterviewService,
        private fb: FormBuilder,
        private router: Router
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

    
    addInterview(interviewDate, feedback, Requisition, Candidate, Interviewers, Stage, Result): void {
        this.interviewService
        .addInterview(interviewDate, feedback, Requisition, Candidate, Interviewers, Stage, Result)
            .subscribe(() => {
                this.router.navigate(['/indexInterview']);
            });
    }

    ngOnInit(): void {
    }
}