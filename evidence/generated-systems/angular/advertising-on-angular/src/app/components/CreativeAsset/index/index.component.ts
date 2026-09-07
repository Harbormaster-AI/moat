
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CreativeAssetService } from '../../../services/CreativeAsset.service';
import { CreativeAsset } from '../../../models/CreativeAsset';

@Component({
    selector: 'app-index-creativeAsset',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCreativeAssetComponent implements OnInit {

    creativeAssets: CreativeAsset[] = [];

    constructor(
        private router: Router,
        private service: CreativeAssetService
) {}

    ngOnInit(): void {
        this.getCreativeAssets();
}

    getCreativeAssets(): void {
        this.service.getCreativeAssets().subscribe((res) => {
        this.creativeAssets = res;
    });
}

    deleteCreativeAsset(id: any): void {
        this.service.deleteCreativeAsset(id)
            .subscribe(() => {
                this.getCreativeAssets();
            });
    }
}