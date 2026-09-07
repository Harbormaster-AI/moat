
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexLegalHoldComponent } from './index.component';
import { LegalHoldService } from '../../../services/LegalHold.service';

describe('IndexLegalHoldComponent', () => {
  let component: IndexLegalHoldComponent;
  let fixture: ComponentFixture<IndexLegalHoldComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexLegalHoldComponent
      ],
      providers: [
        LegalHoldService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexLegalHoldComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});