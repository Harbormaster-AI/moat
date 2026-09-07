
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexWarrantyComponent } from './index.component';
import { WarrantyService } from '../../../services/Warranty.service';

describe('IndexWarrantyComponent', () => {
  let component: IndexWarrantyComponent;
  let fixture: ComponentFixture<IndexWarrantyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexWarrantyComponent
      ],
      providers: [
        WarrantyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexWarrantyComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});