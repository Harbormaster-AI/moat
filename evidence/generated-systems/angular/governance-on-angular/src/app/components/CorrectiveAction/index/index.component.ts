
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CorrectiveActionService } from '../../../services/CorrectiveAction.service';
import { CorrectiveAction } from '../../../models/CorrectiveAction';

@Component({
    selector: 'app-index-correctiveAction',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCorrectiveActionComponent implements OnInit {

    correctiveActions: CorrectiveAction[] = [];

    constructor(
        private router: Router,
        private service: CorrectiveActionService
) {}

    ngOnInit(): void {
        this.getCorrectiveActions();
}

    getCorrectiveActions(): void {
        this.service.getCorrectiveActions().subscribe((res) => {
        this.correctiveActions = res;
    });
}

    deleteCorrectiveAction(id: any): void {
        this.service.deleteCorrectiveAction(id)
            .subscribe(() => {
                this.getCorrectiveActions();
            });
    }
}