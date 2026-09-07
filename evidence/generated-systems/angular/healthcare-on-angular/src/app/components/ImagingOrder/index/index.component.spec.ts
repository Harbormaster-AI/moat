
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexImagingOrderComponent } from './index.component';
import { ImagingOrderService } from '../../../services/ImagingOrder.service';

describe('IndexImagingOrderComponent', () => {
  let component: IndexImagingOrderComponent;
  let fixture: ComponentFixture<IndexImagingOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexImagingOrderComponent
      ],
      providers: [
        ImagingOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexImagingOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});