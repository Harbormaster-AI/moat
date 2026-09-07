import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { IssueService } from '../../../services/Issue.service';
import { SubBaseComponent } from '../../Issue/sub.base.component';


@Component({
    selector: 'app-edit-issue',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditIssueComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Issue';

    issueForm: FormGroup;
    issue: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: IssueService,
        private fb: FormBuilder
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

    
    updateIssue(title, openedDate, closedDate, Risk, Finding, CorrectiveActions, Control, IssueType, Priority, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateIssue(title, openedDate, closedDate, Risk, Finding, CorrectiveActions, Control, IssueType, Priority, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexIssue']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getIssue(params['id']).subscribe(res => {
                this.issue = res;
            });
        });
    }
}