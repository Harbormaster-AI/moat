
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexVerifiedAddressComponent } from './index.component';
import { VerifiedAddressService } from '../../../services/VerifiedAddress.service';

describe('IndexVerifiedAddressComponent', () => {
  let component: IndexVerifiedAddressComponent;
  let fixture: ComponentFixture<IndexVerifiedAddressComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexVerifiedAddressComponent
      ],
      providers: [
        VerifiedAddressService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexVerifiedAddressComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});