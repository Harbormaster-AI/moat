
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AssetService } from '../../../services/Asset.service';
import { Asset } from '../../../models/Asset';

@Component({
    selector: 'app-index-asset',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAssetComponent implements OnInit {

    assets: Asset[] = [];

    constructor(
        private router: Router,
        private service: AssetService
) {}

    ngOnInit(): void {
        this.getAssets();
}

    getAssets(): void {
        this.service.getAssets().subscribe((res) => {
        this.assets = res;
    });
}

    deleteAsset(id: any): void {
        this.service.deleteAsset(id)
            .subscribe(() => {
                this.getAssets();
            });
    }
}