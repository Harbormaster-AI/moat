
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAnalyticsWorkspaceComponent } from './index.component';
import { AnalyticsWorkspaceService } from '../../../services/AnalyticsWorkspace.service';

describe('IndexAnalyticsWorkspaceComponent', () => {
  let component: IndexAnalyticsWorkspaceComponent;
  let fixture: ComponentFixture<IndexAnalyticsWorkspaceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAnalyticsWorkspaceComponent
      ],
      providers: [
        AnalyticsWorkspaceService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAnalyticsWorkspaceComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});