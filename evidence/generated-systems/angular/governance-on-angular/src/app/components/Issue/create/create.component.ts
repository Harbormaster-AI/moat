import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { IssueService } from '../../../services/Issue.service';
import { Issue } from '../../../models/Issue';
import { SubBaseComponent } from '../../Issue/sub.base.component';

@Component({
    selector: 'app-create-issue',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateIssueComponent extends SubBaseComponent implements OnInit {

    title = 'Add Issue';

    issueForm: FormGroup;
    issue: Issue;

    constructor( http: HttpClient,
        private issueService: IssueService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.issueForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      openedDate: ['', Validators.required],
      closedDate: ['', Validators.required],
      Risk: ['', ],
      Finding: ['', ],
      CorrectiveActions: ['', ],
      Control: ['', ],
      IssueType: ['', ],
      Priority: ['', ],
      Status: ['', ]
        });
    }

    
    addIssue(title, openedDate, closedDate, Risk, Finding, CorrectiveActions, Control, IssueType, Priority, Status): void {
        this.issueService
        .addIssue(title, openedDate, closedDate, Risk, Finding, CorrectiveActions, Control, IssueType, Priority, Status)
            .subscribe(() => {
                this.router.navigate(['/indexIssue']);
            });
    }

    ngOnInit(): void {
    }
}