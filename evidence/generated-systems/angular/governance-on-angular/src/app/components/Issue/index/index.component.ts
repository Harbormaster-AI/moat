
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { IssueService } from '../../../services/Issue.service';
import { Issue } from '../../../models/Issue';

@Component({
    selector: 'app-index-issue',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexIssueComponent implements OnInit {

    issues: Issue[] = [];

    constructor(
        private router: Router,
        private service: IssueService
) {}

    ngOnInit(): void {
        this.getIssues();
}

    getIssues(): void {
        this.service.getIssues().subscribe((res) => {
        this.issues = res;
    });
}

    deleteIssue(id: any): void {
        this.service.deleteIssue(id)
            .subscribe(() => {
                this.getIssues();
            });
    }
}