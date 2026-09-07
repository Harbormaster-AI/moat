
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BusinessGlossaryTermService } from '../../../services/BusinessGlossaryTerm.service';
import { BusinessGlossaryTerm } from '../../../models/BusinessGlossaryTerm';

@Component({
    selector: 'app-index-businessGlossaryTerm',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexBusinessGlossaryTermComponent implements OnInit {

    businessGlossaryTerms: BusinessGlossaryTerm[] = [];

    constructor(
        private router: Router,
        private service: BusinessGlossaryTermService
) {}

    ngOnInit(): void {
        this.getBusinessGlossaryTerms();
}

    getBusinessGlossaryTerms(): void {
        this.service.getBusinessGlossaryTerms().subscribe((res) => {
        this.businessGlossaryTerms = res;
    });
}

    deleteBusinessGlossaryTerm(id: any): void {
        this.service.deleteBusinessGlossaryTerm(id)
            .subscribe(() => {
                this.getBusinessGlossaryTerms();
            });
    }
}