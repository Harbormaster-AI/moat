
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCoverageDefinitionComponent } from './index.component';
import { CoverageDefinitionService } from '../../../services/CoverageDefinition.service';

describe('IndexCoverageDefinitionComponent', () => {
  let component: IndexCoverageDefinitionComponent;
  let fixture: ComponentFixture<IndexCoverageDefinitionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCoverageDefinitionComponent
      ],
      providers: [
        CoverageDefinitionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCoverageDefinitionComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});