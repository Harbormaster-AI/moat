
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UnderwriterService } from '../../../services/Underwriter.service';
import { Underwriter } from '../../../models/Underwriter';

@Component({
    selector: 'app-index-underwriter',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexUnderwriterComponent implements OnInit {

    underwriters: Underwriter[] = [];

    constructor(
        private router: Router,
        private service: UnderwriterService
) {}

    ngOnInit(): void {
        this.getUnderwriters();
}

    getUnderwriters(): void {
        this.service.getUnderwriters().subscribe((res) => {
        this.underwriters = res;
    });
}

    deleteUnderwriter(id: any): void {
        this.service.deleteUnderwriter(id)
            .subscribe(() => {
                this.getUnderwriters();
            });
    }
}