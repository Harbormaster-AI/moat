
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TerminationService } from '../../../services/Termination.service';
import { Termination } from '../../../models/Termination';

@Component({
    selector: 'app-index-termination',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTerminationComponent implements OnInit {

    terminations: Termination[] = [];

    constructor(
        private router: Router,
        private service: TerminationService
) {}

    ngOnInit(): void {
        this.getTerminations();
}

    getTerminations(): void {
        this.service.getTerminations().subscribe((res) => {
        this.terminations = res;
    });
}

    deleteTermination(id: any): void {
        this.service.deleteTermination(id)
            .subscribe(() => {
                this.getTerminations();
            });
    }
}