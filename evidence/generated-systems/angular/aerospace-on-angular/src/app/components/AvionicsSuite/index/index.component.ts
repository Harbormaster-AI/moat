
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AvionicsSuiteService } from '../../../services/AvionicsSuite.service';
import { AvionicsSuite } from '../../../models/AvionicsSuite';

@Component({
    selector: 'app-index-avionicsSuite',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAvionicsSuiteComponent implements OnInit {

    avionicsSuites: AvionicsSuite[] = [];

    constructor(
        private router: Router,
        private service: AvionicsSuiteService
) {}

    ngOnInit(): void {
        this.getAvionicsSuites();
}

    getAvionicsSuites(): void {
        this.service.getAvionicsSuites().subscribe((res) => {
        this.avionicsSuites = res;
    });
}

    deleteAvionicsSuite(id: any): void {
        this.service.deleteAvionicsSuite(id)
            .subscribe(() => {
                this.getAvionicsSuites();
            });
    }
}