
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexProductOfferingComponent } from './index.component';
import { ProductOfferingService } from '../../../services/ProductOffering.service';

describe('IndexProductOfferingComponent', () => {
  let component: IndexProductOfferingComponent;
  let fixture: ComponentFixture<IndexProductOfferingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexProductOfferingComponent
      ],
      providers: [
        ProductOfferingService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexProductOfferingComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});