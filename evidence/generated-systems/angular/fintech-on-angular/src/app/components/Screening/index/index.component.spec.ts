
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexScreeningComponent } from './index.component';
import { ScreeningService } from '../../../services/Screening.service';

describe('IndexScreeningComponent', () => {
  let component: IndexScreeningComponent;
  let fixture: ComponentFixture<IndexScreeningComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexScreeningComponent
      ],
      providers: [
        ScreeningService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexScreeningComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});