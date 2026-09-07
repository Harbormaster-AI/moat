import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { GoalService } from '../../../services/Goal.service';
import { Goal } from '../../../models/Goal';
import { SubBaseComponent } from '../../Goal/sub.base.component';

@Component({
    selector: 'app-create-goal',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateGoalComponent extends SubBaseComponent implements OnInit {

    title = 'Add Goal';

    goalForm: FormGroup;
    goal: Goal;

    constructor( http: HttpClient,
        private goalService: GoalService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.goalForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      description: ['', Validators.required],
      targetDate: ['', Validators.required],
      weight: ['', Validators.required],
      Employee: ['', ],
      Cycle: ['', ],
      ParentGoal: ['', ],
      ChildGoals: ['', ],
      Status: ['', ]
        });
    }

    
    addGoal(title, description, targetDate, weight, Employee, Cycle, ParentGoal, ChildGoals, Status): void {
        this.goalService
        .addGoal(title, description, targetDate, weight, Employee, Cycle, ParentGoal, ChildGoals, Status)
            .subscribe(() => {
                this.router.navigate(['/indexGoal']);
            });
    }

    ngOnInit(): void {
    }
}