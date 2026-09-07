
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { QualityRuleService } from '../../../services/QualityRule.service';
import { QualityRule } from '../../../models/QualityRule';

@Component({
    selector: 'app-index-qualityRule',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexQualityRuleComponent implements OnInit {

    qualityRules: QualityRule[] = [];

    constructor(
        private router: Router,
        private service: QualityRuleService
) {}

    ngOnInit(): void {
        this.getQualityRules();
}

    getQualityRules(): void {
        this.service.getQualityRules().subscribe((res) => {
        this.qualityRules = res;
    });
}

    deleteQualityRule(id: any): void {
        this.service.deleteQualityRule(id)
            .subscribe(() => {
                this.getQualityRules();
            });
    }
}