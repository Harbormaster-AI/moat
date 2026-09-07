
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexUnderwriterComponent } from './index.component';
import { UnderwriterService } from '../../../services/Underwriter.service';

describe('IndexUnderwriterComponent', () => {
  let component: IndexUnderwriterComponent;
  let fixture: ComponentFixture<IndexUnderwriterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexUnderwriterComponent
      ],
      providers: [
        UnderwriterService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexUnderwriterComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});