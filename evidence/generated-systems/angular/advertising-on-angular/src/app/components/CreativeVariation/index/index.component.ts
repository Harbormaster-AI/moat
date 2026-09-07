
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CreativeVariationService } from '../../../services/CreativeVariation.service';
import { CreativeVariation } from '../../../models/CreativeVariation';

@Component({
    selector: 'app-index-creativeVariation',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCreativeVariationComponent implements OnInit {

    creativeVariations: CreativeVariation[] = [];

    constructor(
        private router: Router,
        private service: CreativeVariationService
) {}

    ngOnInit(): void {
        this.getCreativeVariations();
}

    getCreativeVariations(): void {
        this.service.getCreativeVariations().subscribe((res) => {
        this.creativeVariations = res;
    });
}

    deleteCreativeVariation(id: any): void {
        this.service.deleteCreativeVariation(id)
            .subscribe(() => {
                this.getCreativeVariations();
            });
    }
}