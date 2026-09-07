
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UnderwritingDecisionService } from '../../../services/UnderwritingDecision.service';
import { UnderwritingDecision } from '../../../models/UnderwritingDecision';

@Component({
    selector: 'app-index-underwritingDecision',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexUnderwritingDecisionComponent implements OnInit {

    underwritingDecisions: UnderwritingDecision[] = [];

    constructor(
        private router: Router,
        private service: UnderwritingDecisionService
) {}

    ngOnInit(): void {
        this.getUnderwritingDecisions();
}

    getUnderwritingDecisions(): void {
        this.service.getUnderwritingDecisions().subscribe((res) => {
        this.underwritingDecisions = res;
    });
}

    deleteUnderwritingDecision(id: any): void {
        this.service.deleteUnderwritingDecision(id)
            .subscribe(() => {
                this.getUnderwritingDecisions();
            });
    }
}