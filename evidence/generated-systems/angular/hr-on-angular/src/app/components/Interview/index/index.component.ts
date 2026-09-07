
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InterviewService } from '../../../services/Interview.service';
import { Interview } from '../../../models/Interview';

@Component({
    selector: 'app-index-interview',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInterviewComponent implements OnInit {

    interviews: Interview[] = [];

    constructor(
        private router: Router,
        private service: InterviewService
) {}

    ngOnInit(): void {
        this.getInterviews();
}

    getInterviews(): void {
        this.service.getInterviews().subscribe((res) => {
        this.interviews = res;
    });
}

    deleteInterview(id: any): void {
        this.service.deleteInterview(id)
            .subscribe(() => {
                this.getInterviews();
            });
    }
}