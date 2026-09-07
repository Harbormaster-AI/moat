
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { GoalService } from '../../../services/Goal.service';
import { Goal } from '../../../models/Goal';

@Component({
    selector: 'app-index-goal',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexGoalComponent implements OnInit {

    goals: Goal[] = [];

    constructor(
        private router: Router,
        private service: GoalService
) {}

    ngOnInit(): void {
        this.getGoals();
}

    getGoals(): void {
        this.service.getGoals().subscribe((res) => {
        this.goals = res;
    });
}

    deleteGoal(id: any): void {
        this.service.deleteGoal(id)
            .subscribe(() => {
                this.getGoals();
            });
    }
}