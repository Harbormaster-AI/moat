
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCoverageComponent } from './index.component';
import { CoverageService } from '../../../services/Coverage.service';

describe('IndexCoverageComponent', () => {
  let component: IndexCoverageComponent;
  let fixture: ComponentFixture<IndexCoverageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCoverageComponent
      ],
      providers: [
        CoverageService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCoverageComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});