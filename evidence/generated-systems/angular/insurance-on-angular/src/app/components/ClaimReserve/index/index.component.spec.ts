
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexClaimReserveComponent } from './index.component';
import { ClaimReserveService } from '../../../services/ClaimReserve.service';

describe('IndexClaimReserveComponent', () => {
  let component: IndexClaimReserveComponent;
  let fixture: ComponentFixture<IndexClaimReserveComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexClaimReserveComponent
      ],
      providers: [
        ClaimReserveService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexClaimReserveComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});