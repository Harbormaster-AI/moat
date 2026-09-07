
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexFXDealComponent } from './index.component';
import { FXDealService } from '../../../services/FXDeal.service';

describe('IndexFXDealComponent', () => {
  let component: IndexFXDealComponent;
  let fixture: ComponentFixture<IndexFXDealComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexFXDealComponent
      ],
      providers: [
        FXDealService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexFXDealComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});