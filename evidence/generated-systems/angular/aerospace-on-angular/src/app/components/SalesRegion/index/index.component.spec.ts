
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexSalesRegionComponent } from './index.component';
import { SalesRegionService } from '../../../services/SalesRegion.service';

describe('IndexSalesRegionComponent', () => {
  let component: IndexSalesRegionComponent;
  let fixture: ComponentFixture<IndexSalesRegionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexSalesRegionComponent
      ],
      providers: [
        SalesRegionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexSalesRegionComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});