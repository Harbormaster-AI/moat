
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DependentService } from '../../../services/Dependent.service';
import { Dependent } from '../../../models/Dependent';

@Component({
    selector: 'app-index-dependent',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDependentComponent implements OnInit {

    dependents: Dependent[] = [];

    constructor(
        private router: Router,
        private service: DependentService
) {}

    ngOnInit(): void {
        this.getDependents();
}

    getDependents(): void {
        this.service.getDependents().subscribe((res) => {
        this.dependents = res;
    });
}

    deleteDependent(id: any): void {
        this.service.deleteDependent(id)
            .subscribe(() => {
                this.getDependents();
            });
    }
}