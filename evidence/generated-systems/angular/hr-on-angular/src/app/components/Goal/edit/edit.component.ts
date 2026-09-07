import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { GoalService } from '../../../services/Goal.service';
import { SubBaseComponent } from '../../Goal/sub.base.component';


@Component({
    selector: 'app-edit-goal',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditGoalComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Goal';

    goalForm: FormGroup;
    goal: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: GoalService,
        private fb: FormBuilder
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

    
    updateGoal(title, description, targetDate, weight, Employee, Cycle, ParentGoal, ChildGoals, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateGoal(title, description, targetDate, weight, Employee, Cycle, ParentGoal, ChildGoals, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexGoal']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getGoal(params['id']).subscribe(res => {
                this.goal = res;
            });
        });
    }
}