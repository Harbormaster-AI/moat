
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCardTokenizationComponent } from './index.component';
import { CardTokenizationService } from '../../../services/CardTokenization.service';

describe('IndexCardTokenizationComponent', () => {
  let component: IndexCardTokenizationComponent;
  let fixture: ComponentFixture<IndexCardTokenizationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCardTokenizationComponent
      ],
      providers: [
        CardTokenizationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCardTokenizationComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});