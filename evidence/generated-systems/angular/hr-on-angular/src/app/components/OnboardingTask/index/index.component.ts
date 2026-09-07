
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { OnboardingTaskService } from '../../../services/OnboardingTask.service';
import { OnboardingTask } from '../../../models/OnboardingTask';

@Component({
    selector: 'app-index-onboardingTask',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexOnboardingTaskComponent implements OnInit {

    onboardingTasks: OnboardingTask[] = [];

    constructor(
        private router: Router,
        private service: OnboardingTaskService
) {}

    ngOnInit(): void {
        this.getOnboardingTasks();
}

    getOnboardingTasks(): void {
        this.service.getOnboardingTasks().subscribe((res) => {
        this.onboardingTasks = res;
    });
}

    deleteOnboardingTask(id: any): void {
        this.service.deleteOnboardingTask(id)
            .subscribe(() => {
                this.getOnboardingTasks();
            });
    }
}