
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ConditionService } from '../../../services/Condition.service';
import { Condition } from '../../../models/Condition';

@Component({
    selector: 'app-index-condition',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexConditionComponent implements OnInit {

    conditions: Condition[] = [];

    constructor(
        private router: Router,
        private service: ConditionService
) {}

    ngOnInit(): void {
        this.getConditions();
}

    getConditions(): void {
        this.service.getConditions().subscribe((res) => {
        this.conditions = res;
    });
}

    deleteCondition(id: any): void {
        this.service.deleteCondition(id)
            .subscribe(() => {
                this.getConditions();
            });
    }
}