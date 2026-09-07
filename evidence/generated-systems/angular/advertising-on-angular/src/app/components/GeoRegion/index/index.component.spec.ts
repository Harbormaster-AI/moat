
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexGeoRegionComponent } from './index.component';
import { GeoRegionService } from '../../../services/GeoRegion.service';

describe('IndexGeoRegionComponent', () => {
  let component: IndexGeoRegionComponent;
  let fixture: ComponentFixture<IndexGeoRegionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexGeoRegionComponent
      ],
      providers: [
        GeoRegionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexGeoRegionComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});